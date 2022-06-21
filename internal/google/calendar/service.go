package calendar

import (
	"context"
	"fmt"
	"time"

	"iot/internal/config"
	"iot/pkg/automation"
	"iot/pkg/google"
	"iot/pkg/mqtt"
)

type CalendarService struct {
	calendarGoogle    *google.CalendarGoogle
	automationService *automation.AutomationService
	mqttClient        *mqtt.MqttService
	event             config.MqttEvent
}

func NewCalendarService(cg *google.CalendarGoogle, mqttClient *mqtt.MqttService, event config.MqttEvent) *CalendarService {
	return &CalendarService{
		calendarGoogle:    cg,
		automationService: automation.NewAutomationService(),
		mqttClient:        mqttClient,
	}
}

func (cs *CalendarService) GetEvents() {
	ctx := context.Background()
	for {
		time.Sleep(10 * time.Second)
		cs.GetEvent(ctx)
	}
}

func (cs *CalendarService) GetEvent(ctx context.Context) {
	fmt.Println("Entrou nos events ")
	event, err := cs.calendarGoogle.Get(ctx)
	if err != nil {
		cs.powerOff(ctx)
	}
	fmt.Println("start", event.End.DateTime)
	fmt.Println("start", event.End.TimeZone)

	loc, err := time.LoadLocation(event.Start.TimeZone)
	if err != nil {
		fmt.Println(err)
		cs.powerOff(ctx)
	}

	endTime, err := time.ParseInLocation(time.RFC3339, event.End.DateTime, loc)
	if err != nil {
		fmt.Println(err)
		cs.powerOff(ctx)
	}
	startTime, err := time.ParseInLocation(time.RFC3339, event.Start.DateTime, loc)
	if err != nil {
		fmt.Println(err)
		cs.powerOff(ctx)
	}
	now := time.Now().In(loc)

	if now.After(startTime) && now.Before(endTime) {
		fmt.Println("Existe")

		cs.mqttClient.Publish(cs.event.Calendar, event.Summary)

		cs.powerOn(ctx)
	} else {
		cs.powerOff(ctx)
	}
	fmt.Println(startTime, endTime, now)
}

func (cs *CalendarService) powerOff(ctx context.Context) {
	cs.automationService.Blink(ctx, "no")
}

func (cs *CalendarService) powerOn(ctx context.Context) {
	cs.automationService.Blink(ctx, "yes")
}
