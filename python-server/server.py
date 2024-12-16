import grpc
from concurrent import futures
import qa_pb2
import qa_pb2_grpc
import os
from model import response_from_llm


class QAServicer(qa_pb2_grpc.QAServiceServicer):
    def AskQuestion(self, request, context):
        question = request.question.lower()
        llm_answer = response_from_llm(question)
        return qa_pb2.QuestionResponse(answer=llm_answer)


def serve():
    # 환경 변수 설정
    os.environ["PYTHONIOENCODING"] = "utf-8"

    # gRPC 서버 초기화
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    qa_pb2_grpc.add_QAServiceServicer_to_server(QAServicer(), server)
    server.add_insecure_port('[::]:50051')

    print("Python 서버가 50051 포트에서 시작되었습니다.")

    try:
        server.start()
        server.wait_for_termination()
    except KeyboardInterrupt:
        print("\nPython 서버가 사용자 요청으로 종료되었습니다.")
    finally:
        server.stop(0)


if __name__ == '__main__':
    serve()
