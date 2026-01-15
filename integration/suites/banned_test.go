package suites

import (
	"fmt"
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

// If there's a DONKEY in the system and we ask for SHREK, what happens?
func TestBannedPlayerDoesNotMatch(t *testing.T) {
	// setting up a dynamoService to have access to dynamodb
	dynamoService, _ := localstack.NewDynamoClient()

	// adding "DONKEY" into the banned players table
	helpers.AddBannedPlayer("DONKEY", dynamoService)

	// remove "DONKEY" from banned players table - once everything else is finished
	defer helpers.RemoveBannedPlayer("DONKEY", dynamoService)

	// run the banned name check and store the result in output
	output := helpers.RunBannedCheck("SHREK")

	// assert that the output will be false (SHREK is not in the banned players table)
	assert.Equal(t, output, "false")
}

// If theres no players in the system, what happens when the app runs?
func TestNoPlayers(t *testing.T) {
	output := helpers.RunBannedCheck("SHREK")

	assert.Equal(t, output, "false")
}

// If multiple banned players are in the system, what happens when the app tries to query each of them?
// And then query a player that then doesnt exist while it's in this state?
func TestMultipleBannedPlayers(t *testing.T) {
	dynamoService, _ := localstack.NewDynamoClient()

	helpers.AddBannedPlayer("DONKEY", dynamoService)
	defer helpers.RemoveBannedPlayer("DONKEY", dynamoService)

	helpers.AddBannedPlayer("SHREK", dynamoService)
	defer helpers.RemoveBannedPlayer("SHREK", dynamoService)

	helpers.AddBannedPlayer("FIONA", dynamoService)
	defer helpers.RemoveBannedPlayer("FIONA", dynamoService)

	outputOne := helpers.RunBannedCheck("DONKEY")
	outputTwo := helpers.RunBannedCheck("SHREK")
	outputThree := helpers.RunBannedCheck("FIONA")
	outputFour := helpers.RunBannedCheck("LORDFARQUAAD")
	fmt.Println("im different")

	// assert that the output will be true for all except LORDFARQUAAD which is not in the banned players table
	assert.Equal(t, outputOne, "true")
	assert.Equal(t, outputTwo, "true")
	assert.Equal(t, outputThree, "true")
	assert.Equal(t, outputFour, "false")
}
