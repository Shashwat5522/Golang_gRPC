package main
import(
	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
	"log"
	"fmt"
)
func(s *Server) GreetManyTimes(request *pb.GreetRequest,stream pb.GreetService_GreetManyTimesServer)error{
	log.Printf("greet many times function invoked")
	
	for i:=0;i<10;i++{
		res:=fmt.Sprintf("Hello %s,number %d",request.FirstName,i)
		stream.Send(&pb.GreetResponse{
			Result:res,
		})
	}
	return nil
}