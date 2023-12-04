package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aleksi1212/cobra-http-request-cli/internal/utils"
	"github.com/spf13/cobra"
)

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Makes an HTTP request",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		method, _ := cmd.Flags().GetString("method")
		url, _ := cmd.Flags().GetString("url")
		headers, _ := cmd.Flags().GetString("headers")
		filePath, _ := cmd.Flags().GetString("filepath")

		var headerArray utils.HeadersArray

		if url == "" {
			fmt.Println("please provide a url\n\nFlags:\n	--url, -u")
			return
		}
		if headers != "" {
			err := headerArray.Set(headers)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		response := makeRequest(
			method,
			url,
			headerArray,
			filePath,
		)
		fmt.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(requestCmd)

	requestCmd.Flags().StringP("method", "m", "GET", "insert an HTTP method")
	requestCmd.Flags().StringP("url", "u", "", "insert a url *")
	requestCmd.Flags().StringP("headers", "H", "", "insert HTTP headers")
	requestCmd.Flags().StringP("body", "D", "", "insert body data")
	requestCmd.Flags().StringP("filepath", "P", "", "insert file path to where to write response body")
}

func makeRequest(method string, url string, headers utils.HeadersArray, filePath string) string {
	req, reqError := http.NewRequest(method, url, nil)
	if reqError != nil {
		return fmt.Sprintf("Error creating request: %v", reqError)
	}

	if len(headers) > 0 {
		for _, value := range headers {
			req.Header.Set(value[0], value[1])
		}
	}

	client := &http.Client{}
	resp, respError := client.Do(req)
	if respError != nil {
		return fmt.Sprintf("Error performing request: %v", respError)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	body, readError := io.ReadAll(resp.Body)
	if readError != nil {
		return fmt.Sprintf("Error reading response body: %v", readError)
	}

	if filePath != "" {
		return utils.FileWriter(filePath, contentType, body)
	}

	return string(body)
}
