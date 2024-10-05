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

	// Run our service in it's entirety

	// Our service will use the amazon SDK to try to retrieve the data it needs
	// However, we are forcibly making the SDK talk to a container running on our
	// own machine. It will never reach REAL AWS, and instead will retrieve the data we
	// just inserted earlier
	output := helpers.RunBannedCheck("SHREK")

	// Our service just ran and it thinks that it was talking to a real AWS DynamoDB
	// Check the output! Did our service detect the banned player?
	assert.Equal(t, output, "true")

	// Finally, reset localstack to it's original state for future tests
	helpers.RemoveBannedPlayer("SHREK", dynamoService)
}

// And now we can simulate a bunch of scenarios that dynamoDB could respond in!
// And we can confirm our service handles it all correctly!
func TestBannedPlayerDoesNotMatch(t *testing.T) {
	dynamoService, _ := localstack.NewDynamoClient()

	helpers.AddBannedPlayer("SHREK", dynamoService)

	assert.Equal(t, helpers.RunBannedCheck("SHREK"), "false")
	helpers.RemoveBannedPlayer("SHREK", dynamoService)
}

func TestNoBannedPlayers(t *testing.T) {
	assert.Equal(t, helpers.RunBannedCheck("SHREK"), "false")
}

func TestMultipleBannedPlayers(t *testing.T) {
	dynamoService, _ := localstack.NewDynamoClient()

	helpers.AddBannedPlayer("SHREK", dynamoService)
	helpers.AddBannedPlayer("DONKEY", dynamoService)
	helpers.AddBannedPlayer("FIONA", dynamoService)

	assert.Equal(t, helpers.RunBannedCheck("SHREK"), "true")
	assert.Equal(t, helpers.RunBannedCheck("DONKEY"), "true")
	assert.Equal(t, helpers.RunBannedCheck("FIONA"), "true")

	assert.Equal(t, helpers.RunBannedCheck("DRAGON"), "false")
	assert.Equal(t, helpers.RunBannedCheck("FARQUAAD"), "false")

	helpers.RemoveBannedPlayer("SHREK", dynamoService)
	helpers.RemoveBannedPlayer("DONKEY", dynamoService)
	helpers.RemoveBannedPlayer("FIONA", dynamoService)
}
