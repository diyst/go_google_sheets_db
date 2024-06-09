package sheets

import "google.golang.org/api/sheets/v4"

func (c *clientModel) Read(spreadsheetId, readRange string) ([][]interface{}, error) {
	resp, err := c.Service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}

func (c *clientModel) Write(spreadsheetId, writeRange string, values [][]interface{}) error {
	vr := &sheets.ValueRange{
		Values: values,
	}
	_, err := c.Service.Spreadsheets.Values.Append(spreadsheetId, writeRange, vr).ValueInputOption("RAW").Do()

	c.Service.Spreadsheets.Values.BatchGet(spreadsheetId).Do()

	return err
}
