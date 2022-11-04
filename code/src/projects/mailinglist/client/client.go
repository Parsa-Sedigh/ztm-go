package main

import (
	"context"
	"github.com/alexflint/go-arg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "mailinglist/proto"
	"time"
)

// a utility function to log the responses from the grpc server:
func logResponse(res *pb.EmailResponse, err error) {
	if err != nil {
		// these spaces are for readability
		log.Fatalf("  error: %v", err)
	}

	if res.EmailEntry == nil {
		log.Printf("  email not found")
	} else {
		log.Printf("  response: %v", res.EmailEntry)
	}
}

func createEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Println("create email")

	/* Our grpc server requires that we have a context with the req. We create one with WithTimeout() function and this function will automatically cancel the
	req after a certain amount of time. In this specific scenario, we're using time.Second , so the req has 1 second to complete.*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	/* If the req takes less than 1 second and the function completes, since we defined 1 second as the maximum in the line above, we're gonna run the
	cancel() function which will free up any resources. That way, the resources aren't being used for an entire second.*/
	defer cancel()

	/* When we ran protoc toll on the protocol buffers file, it created both server and client functions that we're able to utilize. Here, we're using the client
	function named CreateEmail() to make a req to the server.*/
	res, err := client.CreateEmail(ctx, &pb.CreateEmailRequest{EmailAddr: addr})
	logResponse(res, err)

	return res.EmailEntry
}

func getEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Println("get email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	res, err := client.GetEmail(ctx, &pb.GetEmailRequest{EmailAddr: addr})
	logResponse(res, err)

	return res.EmailEntry
}

func getEmailBatch(client pb.MailingListServiceClient, count int, page int) {
	log.Println("get email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	res, err := client.GetBatchEmail(ctx, &pb.GetEmailBatchRequest{Count: int32(count), Page: int32(page)})

	// we can't use logResponse() for this req, because it's functionality is a bit different, so we write code for logging the response here manually
	if err != nil {
		log.Fatalf("  error: %v", err)
	}

	log.Println("response:")

	for i := 0; i < len(res.EmailEntries); i++ {
		log.Printf("  item [%v of %v]: %s", i+1, len(res.EmailEntries), res.EmailEntries[i])
	}
}

func updateEmail(client pb.MailingListServiceClient, entry pb.EmailEntry) *pb.EmailEntry {
	log.Println("update email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	res, err := client.UpdateEmail(ctx, &pb.UpdateEmailRequest{EmailEntry: &entry})
	logResponse(res, err)

	return res.EmailEntry
}

func deleteEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Println("delete email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	res, err := client.DeleteEmail(ctx, &pb.DeleteEmailRequest{EmailAddr: addr})
	logResponse(res, err)

	return res.EmailEntry
}

// command line arguments:
var args struct {
	GrpcAddr string `arg:"env:MAILINGLIST_GRPC_ADDR"`
}

func main() {
	arg.MustParse(&args)

	// set a default for command line argument:
	if args.GrpcAddr == "" {
		args.GrpcAddr = ":8081"
	}

	// connect to the server:
	/* insecure.NewCredentials(): Means we're connecting with an insecure connection, no encryption, anything like that, no authorization, no username required.
	This microservice is meant to be running on your backend and only communicating with pre-authorized services, so we're not concerned with credentials.*/
	conn, err := grpc.Dial(args.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// close the connection after the application terminates:
	defer conn.Close()

	/* When we use NewMailingListServiceClient() function, we're telling the existing grpc connection that this client is associated with the RPC messages that
	we defined within our mailing list service. This will allow us to send reqs to the server.*/
	client := pb.NewMailingListServiceClient(conn)

	/////// SEND REQUESTS //////
	newEmail := createEmail(client, "9999999@999.999")
	newEmail.ConfirmedAt = 10000

	updateEmail(client, *newEmail)
	deleteEmail(client, newEmail.Email)
	getEmailBatch(client, 3, 1)
	getEmailBatch(client, 3, 2)
	getEmailBatch(client, 3, 3)
}
