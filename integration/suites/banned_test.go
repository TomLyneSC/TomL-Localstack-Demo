package suites

import (
	helpers "localstack-demo/integration"
	"localstack-demo/src/localstack"
	"testing"

	"github.com/go-playground/assert"
)

func TestBannedPlayerMatches(t *testing.T) {
	// Setup localstack connection
	dynamoService, _ := localstack.NewDynamoClient()

	// Create a banned player in our "local" DynamoDB
	helpers.AddBannedPlayer("SHREK", dynamoService)

	// At end of test, reset localstack to it's original state for future tests
	defer helpers.RemoveBannedPlayer("SHREK", dynamoService)

	// Run our service in it's entirety

	// Our service will use the amazon SDK to try to retrieve the data it needs
	// However, we are forcibly making the SDK talk to a container running on our
	// own machine. It will never reach REAL AWS, and instead will retrieve the data we
	// just inserted earlier
	output := helpers.RunBannedCheck("SHREK")

	// Our service just ran and it thinks that it was talking to a real AWS DynamoDB
	// Check the output! Did our service detect the banned player?
	assert.Equal(t, output, "true")
}

// And now we can simulate a bunch of scenarios that dynamoDB could respond in!
// And we can confirm our service handles it all correctly and produces the correct output!
func TestBannedPlayerDoesNotMatch(t *testing.T) {
}

func TestNoPlayers(t *testing.T) {
}

func TestMultipleBannedPlayers(t *testing.T) {
}
