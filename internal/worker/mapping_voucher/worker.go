package mapping_voucher

import (
	"context"
	"encoding/json"
	"github.com/ChungKiet/cake-interview/adapter"
	"github.com/ChungKiet/cake-interview/internal/biz"
	log2 "github.com/go-kratos/kratos/v2/log"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	// Init Use Case
	log *log2.Helper

	nc    nats.JetStreamContext
	topic string

	userVoucherRepo     biz.UserVoucherRepo
	campaignRepo        biz.CampaignRepo
	voucherCampaignRepo biz.VoucherCampaignRepo
}

func NewWorker(nc nats.JetStreamContext, topic string, userVoucherRepo biz.UserVoucherRepo, campaignRepo biz.CampaignRepo, voucherCampaignRepo biz.VoucherCampaignRepo) *Worker {
	return &Worker{
		log: log2.NewHelper(log2.DefaultLogger),

		nc:    nc,
		topic: topic,

		userVoucherRepo:     userVoucherRepo,
		campaignRepo:        campaignRepo,
		voucherCampaignRepo: voucherCampaignRepo,
	}
}

func (w *Worker) Run() {
	// Initialize NATS connection

	// Initialize context and wait group
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Handle OS signals to gracefully shutdown
	handleSignals(cancel, &wg)

	// Subscribe to Nats channel
	msgChan := make(chan *nats.Msg, 100)
	sub, err := w.nc.ChanSubscribe(w.topic, msgChan, []nats.SubOpt{nats.OrderedConsumer()}...)
	if err != nil {
		w.log.Error("Subscribe failed", "err", err)
		panic(err)
	}

	// Wait for messages in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-msgChan:
				w.consumerMsg(msg)
			}
		}
	}()

	// Run the service until a signal is received
	wg.Wait()
	sub.Unsubscribe()
	close(msgChan)
}

type Signal struct {
	SignalId     int     `json:"signal_id"`
	Indicator    string  `json:"indicator"`
	Source       string  `json:"source"`
	SignalStatus string  `json:"signal_status"`
	IsUpdate     bool    `json:"is_update"`
	Symbol       string  `json:"symbol"`
	TimeFrame    string  `json:"time_frame"`
	Signal       string  `json:"signal"`
	CloseTime    int64   `json:"close_time"`
	ClosePrice   float32 `json:"close_price"`
	Type         string  `json:"type"` // close / realtime
	SourceDown   string  `json:"source_down"`
	IsTest       bool    `json:"isTest"`
}

// Publish a message to NATS
func publishToNATS(nc *nats.Conn, subject string, message string) {
	err := nc.Publish(subject, []byte(message))
	if err != nil {
		log.Printf("Error publishing to NATS: %v", err)
	}
}

// Handle OS signals for graceful shutdown
func handleSignals(cancel context.CancelFunc, wg *sync.WaitGroup) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("Received signal: %v", sig)
		cancel()
		wg.Wait() // Wait for goroutines to finish before exiting
		os.Exit(0)
	}()
}

// Consumer message
func (w *Worker) consumerMsg(msg *nats.Msg) error {
	if msg == nil {
		return nil
	}

	if len(msg.Data) == 0 {
		return nil
	}

	// parse data
	var event adapter.UserLoginEventData
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		return err
	}

	// TODO: config timeout
	timeOut := 5 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeOut) // don't use context cancel
	errTx := w.userVoucherRepo.WithTx(ctx, func(ctx context.Context) error {
		// finding the campaign not full slot
		campaign, err := w.campaignRepo.GetCampaignNotFull(ctx)
		if err != nil {
			return err
		}

		// find voucher of this campaign
		voucherCampaigns, err := w.voucherCampaignRepo.GetVoucherCampaign(ctx, int64(campaign.ID))
		if err != nil {
			return err
		}

		for _, item := range voucherCampaigns {
			// create voucher campaign
			err = w.userVoucherRepo.CreateUserVoucher(ctx, &biz.CreateUserVoucherRequest{
				VoucherCampaignID: int64(item.ID),
				UserID:            event.UserID,
			})

			if err != nil {
				return err
			}
		}

		// increase campaign slot

		currentSlot := campaign.CurrentSlot + 1
		err = w.campaignRepo.IncreaseCampaignSlot(ctx, &biz.IncreaseCampaignSlotRequest{
			CampaignID:  int64(campaign.ID),
			CurrentSlot: currentSlot,
			IsFull:      currentSlot >= campaign.TotalSlot,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if errTx != nil {
		w.log.Error("Create user_voucher failed", "err", errTx.Error())
		return errTx
	}

	// TODO query by batcher db to check alert

	// TODO get by batcher

	return nil
}
