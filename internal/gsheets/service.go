package gsheets

import "google.golang.org/api/sheets/v4"

func (c *ClientModel) Read(spreadsheetId, readRange string) ([][]interface{}, error) {
	resp, err := c.Service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}

func (c *ClientModel) Write(spreadsheetId, writeRange string, values [][]interface{}) error {
	vr := &sheets.ValueRange{
		Values: values,
	}
	_, err := c.Service.Spreadsheets.Values.Append(spreadsheetId, writeRange, vr).ValueInputOption("RAW").Do()

	c.Service.Spreadsheets.Values.BatchGet(spreadsheetId).Do()

	return err
}

func (c *ClientModel) BatchUpdate(request sheets.Request) (sheets.BatchUpdateSpreadsheetResponse, error) {
	response, err := Client.Service.Spreadsheets.BatchUpdate(Client.SheetID, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			&request,
		},
	}).Do()

	if err != nil {
		return sheets.BatchUpdateSpreadsheetResponse{}, err
	}

	return *response, nil
}
