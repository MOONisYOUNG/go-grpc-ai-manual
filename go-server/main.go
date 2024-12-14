package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	pb "go-grpc-ai-manual/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedQAServiceServer
	qaClient pb.QAServiceClient
}

func main() {
	// gRPC 서버 설정
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("서버 리스닝 실패: %v", err)
	}

	s := grpc.NewServer()

	// Python QA 서버와 통신할 클라이언트 설정
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Python 서버 연결 실패: %v", err)
	}
	defer conn.Close()

	qaClient := pb.NewQAServiceClient(conn)

	serverInstance := &server{
		qaClient: qaClient,
	}
	pb.RegisterQAServiceServer(s, serverInstance)

	// 서버 시작 (비동기)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("서버 실행 실패: %v", err)
		}
	}()

	// CLI 인터페이스 시작
	startCLI(qaClient)
}

func startCLI(client pb.QAServiceClient) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("💧[ SK 매직 ] 에코 미니 정수기💦 사용법 Q&A 서비스. 'exit'를 입력하면 종료됩니다.")

	for {
		fmt.Print("\n🤗 질문을 입력하세요: ")
		input, err := reader.ReadString('\n')
		input = strings.ToValidUTF8(input, "")
		if err != nil {
			log.Printf("입력 오류: %v", err)
			continue
		}

		// 개행 문자 제거
		input = strings.TrimSpace(input)

		// 종료 조건
		if input == "exit" {
			fmt.Println("Q&A CLI를 종료합니다.")
			break
		}

		// 빈 입력 무시
		if input == "" {
			continue
		}

		// Python 서버에 질문 보내기
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		resp, err := client.AskQuestion(ctx, &pb.QuestionRequest{Question: input})
		if err != nil {
			log.Printf("질문 실패: %v", err)
		} else {
			fmt.Printf("답변: %s\n", resp.Answer)
		}
		cancel()
	}
}
