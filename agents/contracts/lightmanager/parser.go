package lightmanager

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// Parser parses events from the light inbox contract.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseDisputeOpened parses a DisputeOpened event
	ParseDisputeOpened(log ethTypes.Log) (_ *LightManagerDisputeOpened, err error)
	// ParseRootUpdated parses a RootUpdated event
	ParseRootUpdated(log ethTypes.Log) (_ *[32]byte, err error)
}

type parserImpl struct {
	// filterer is the parser filterer we use to parse events
	filterer *LightManagerFilterer
}

// NewParser creates a new parser for the light manager contract.
func NewParser(lightManagerAddress common.Address) (Parser, error) {
	parser, err := NewLightManagerFilterer(lightManagerAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", LightManagerFilterer{}, err)
	}

	return &parserImpl{filterer: parser}, nil
}

func (p parserImpl) ParseDisputeOpened(log ethTypes.Log) (_ *LightManagerDisputeOpened, err error) {
	lightManagerDisputeOpened, err := p.filterer.ParseDisputeOpened(log)
	if err != nil {
		return nil, fmt.Errorf("could not parse status updated: %w", err)
	}

	return lightManagerDisputeOpened, nil
}

func (p parserImpl) ParseRootUpdated(log ethTypes.Log) (_ *[32]byte, err error) {
	lightManagerRootUpdated, err := p.filterer.ParseRootUpdated(log)
	if err != nil {
		return nil, fmt.Errorf("could not parse status updated: %w", err)
	}

	return &lightManagerRootUpdated.NewRoot, nil
}

func (p parserImpl) EventType(log ethTypes.Log) (_ EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := eventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}

		return *eventType, true
	}
	// return an unknown event to avoid cases where user failed to check the event type
	return EventType(len(topicMap()) + 2), false
}

// EventType is the type of the light manager event
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// DisputeOpenedEvent is an DisputeOpened event.
	DisputeOpenedEvent EventType = iota
	// RootUpdatedEvent is an RootUpdated event.
	RootUpdatedEvent
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
