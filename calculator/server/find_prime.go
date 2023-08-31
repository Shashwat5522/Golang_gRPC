package main

import (
	"log"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
)

func (s *Server) FindPrime(request *pb.CalculatorRequest, stream pb.CalculatorService_FindPrimeServer) error {
	log.Printf("findprime function invoked from server side")

	ans:=findPrimeFactors(int(request.Number))
	for i:=0;i<len(ans);i++{
		res:=ans[i]
		stream.Send(&pb.CalculatorResponse{
			Result:res,
		})
	}
	return nil

}

func findPrimeFactors(number int)[]int32{
	ans:=[]int32{}
	var k int=2
	for number>1{
		if number%k==0{
				ans=append(ans,int32(k))
				number=number/k
		}else{
			k=k+1
		}
		


	}
	return ans
}
