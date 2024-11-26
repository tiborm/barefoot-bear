package fetch

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockPoster is a mock implementation of the Poster interface
type MockPoster struct {
    mock.Mock
}

func (m *MockPoster) Post(url, contentType string, body io.Reader) (*http.Response, error) {
    args := m.Called(url, contentType, body)
    return args.Get(0).(*http.Response), args.Error(1)
}

// MockSleeper is a mock implementation of the Sleeper interface
type MockSleeper struct {
    mock.Mock
}

func (m *MockSleeper) Sleep(duration time.Duration) {
    m.Called(duration)
}

func TestFetch(t *testing.T) {
    mockPoster := new(MockPoster)
    mockSleeper := new(MockSleeper)

    // Create a sample response body
    responseBody := `{"results":[{"items":[{"product":{"id":"123"}},{"product":{"id":"456"}}]}]}`
    mockResponse := &http.Response{
        StatusCode: http.StatusOK,
        Body:       io.NopCloser(bytes.NewBufferString(responseBody)),
    }

    // Set up expectations
    mockPoster.On("Post", "http://example.com", "application/json", mock.Anything).Return(mockResponse, nil)
    mockSleeper.On("Sleep", 2*time.Second).Return()

    // Create an instance of ProductFetcher
    pf := NewProductFetcher(mockSleeper, mockPoster)

    // Define FetchParams
    params := FetchParams{
        URL:            "http://example.com",
        PostPayload:    json.RawMessage(`{"searchParameters":{"input":""}}`),
        FetchSleepTime: 2,
    }

    // Call the Fetch function
    result, err := pf.Fetch("test-id", params)

    // Assert no error
    assert.NoError(t, err)

    // Assert the expected result
    expectedResult := []byte(responseBody)
    assert.Equal(t, expectedResult, result)

    // Assert expectations
    mockPoster.AssertExpectations(t)
    mockSleeper.AssertExpectations(t)
}

func TestGetIDs(t *testing.T) {
    // Sample JSON response
    rawJson := []byte(`{
        "results": [
            {
                "items": [
                    {
                        "product": {
                            "id": "123"
                        }
                    },
                    {
                        "product": {
                            "id": "456"
                        }
                    }
                ]
            }
        ]
    }`)

    // Create an instance of ProductFetcher
    pf := ProductFetcher{}

    // Call the GetIDs function
    ids, err := pf.GetIDs(rawJson)

    // Assert no error
    assert.NoError(t, err)

    // Assert the expected IDs
    expectedIDs := []string{"123", "456"}
    assert.Equal(t, expectedIDs, ids)
}

func TestGetIDs_EmptyResults(t *testing.T) {
    // Sample JSON response with empty results
    rawJson := []byte(`{
        "results": []
    }`)

    // Create an instance of ProductFetcher
    pf := ProductFetcher{}

    // Call the GetIDs function
    ids, err := pf.GetIDs(rawJson)

    // Assert no error
    assert.NoError(t, err)

    // Assert the expected IDs (empty slice)
    expectedIDs := []string{}
    assert.Equal(t, expectedIDs, ids)
}

func TestGetIDs_InvalidJson(t *testing.T) {
    // Invalid JSON response
    rawJson := []byte(`{invalid json}`)

    // Create an instance of ProductFetcher
    pf := ProductFetcher{}

    // Call the GetIDs function
    _, err := pf.GetIDs(rawJson)

    // Assert an error
    assert.Error(t, err)
}