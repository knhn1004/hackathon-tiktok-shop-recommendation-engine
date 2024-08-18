import grpc
from concurrent import futures

from .recommendation_pb2_grpc import add_RecommendationServiceServicer_to_server
from .servicers.recommendation import RecommendationService


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_RecommendationServiceServicer_to_server(RecommendationService(), server)
    port = 50051
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    print(f"Server started on port {port}")
    server.wait_for_termination()


def main():
    serve()


if __name__ == '__main__':
    main()