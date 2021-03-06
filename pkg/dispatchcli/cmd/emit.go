///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	eventTypes "github.com/vmware/dispatch/pkg/events"
	"golang.org/x/net/context"

	"github.com/vmware/dispatch/pkg/dispatchcli/i18n"
	"github.com/vmware/dispatch/pkg/event-manager/gen/client/events"
	models "github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

var (
	emitLong = i18n.T(`Emit an event.`)

	// TODO: Add examples
	emitExample = i18n.T(``)

	emitEventDataBinary   = false
	emitEventData         = ""
	emitEventDataFromFile = ""
	emitEventNamespace    = "dispatchframework.io"
	emitEventSourceID     = "dispatch"
	emitEventSourceType   = "dispatch"
	emitEventID           = ""
	emitEventContentType  = ""
	emitEventTypeVersion  = ""
	emitEventSchemaURL    = ""
)

// NewCmdEmit creates a command to emit a dispatch event.
func NewCmdEmit(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "emit EVENT_TYPE [--data PAYLOAD]|[--data-from-file PATH] [--source-id ID] [--source-type TYPE] [--event-type-version VERSION] [--content-type CONTENT_TYPE] [--event-id EVENT_ID]",
		Short:   i18n.T("Emit a dispatch event"),
		Long:    emitLong,
		Example: emitExample,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := runEmit(out, errOut, cmd, args)
			CheckErr(err)
		},
	}
	cmd.Flags().StringVar(&emitEventTypeVersion, "event-type-version", "", "Event type version.")
	cmd.Flags().StringVar(&emitEventNamespace, "namespace", emitEventNamespace, "Event namespace.")
	cmd.Flags().StringVar(&emitEventSourceID, "source-id", emitEventSourceID, "Event source ID.")
	cmd.Flags().StringVar(&emitEventSourceType, "source-type", emitEventSourceType, "Event source type.")
	cmd.Flags().StringVar(&emitEventID, "event-id", "", "Event ID. If empty, will be autogenerated.")
	cmd.Flags().StringVar(&emitEventContentType, "content-type", "", "Content type of the payload. Defaults to application/json if --data is used. For --data-from-file, it will try to detect the type.")
	cmd.Flags().StringVar(&emitEventData, "data", "", "Event data payload. Mutually exclusive with --data-from-file")
	cmd.Flags().StringVar(&emitEventDataFromFile, "data-from-file", "", "Read event payload from the given file path. Mutually exclusive with --data")
	cmd.Flags().BoolVar(&emitEventDataBinary, "binary", false, "Sets payload as binary (will be base64 encoded). If not specified, automatic detection will be attempted.")
	cmd.Flags().StringVar(&emitEventSchemaURL, "event-schema-url", "", "Event Schema URL.")
	return cmd
}

func runEmit(out, errOut io.Writer, cmd *cobra.Command, args []string) error {

	if emitEventID == "" {
		emitEventID = uuid.NewV4().String()
	}

	data, contentType, err := getData()
	if err != nil {
		fmt.Fprintf(errOut, "Error reading event data: %s", err)
		return err
	}

	emission := &models.Emission{
		Event: &models.CloudEvent{
			CloudEventsVersion: swag.String(eventTypes.CloudEventsVersion),
			ContentType:        contentType,
			Data:               data,
			EventID:            &emitEventID,
			EventTime:          strfmt.DateTime(time.Now()),
			EventType:          &args[0],
			EventTypeVersion:   emitEventTypeVersion,
			Namespace:          &emitEventNamespace,
			SchemaURL:          emitEventSchemaURL,
			SourceID:           &emitEventSourceID,
			SourceType:         &emitEventSourceType,
		},
	}

	params := &events.EmitEventParams{
		Context: context.Background(),
		Body:    emission,
	}
	client := eventManagerClient()
	_, err = client.Events.EmitEvent(params, GetAuthInfoWriter())
	if err != nil {
		return formatAPIError(err, params)
	}
	fmt.Fprintln(out, "event emitted")
	return nil
}

func getData() (data string, contentType string, err error) {
	if emitEventData == "" && emitEventDataFromFile == "" {
		return "", "", nil
	}
	if emitEventData != "" && emitEventDataFromFile != "" {
		return "", "", errors.New("--data and --data-from-file are mutually exclusive, specify only one of them")
	}
	if emitEventData != "" {
		// We assume JSON
		return emitEventData, "application/json", nil
	}

	buf, err := ioutil.ReadFile(emitEventDataFromFile)
	if err != nil {
		return "", "", err
	}

	contentType = http.DetectContentType(buf)
	dataString := ""
	if emitEventDataBinary || !strings.HasPrefix(contentType, "text") {
		dataString = base64.StdEncoding.EncodeToString(buf)
	} else {
		dataString = string(buf)
	}
	if emitEventContentType != "" {
		// We prefer user-supplied content-type
		contentType = emitEventContentType
	}

	return dataString, contentType, nil

}
