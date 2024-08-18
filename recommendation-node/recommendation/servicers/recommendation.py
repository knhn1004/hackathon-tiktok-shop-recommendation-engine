import grpc
from ..recommendation_pb2_grpc import RecommendationServiceServicer
from ..recommendation_pb2 import RecommendationRequest, RecommendationResponse

class RecommendationService(RecommendationServiceServicer):
    def GetRecommendations(self, request: RecommendationRequest, context: grpc.ServicerContext) -> RecommendationResponse:
        # TODO: Implement recommendation logic based on user_id and article_id
        # For now, return dummy product IDs
        return RecommendationResponse(product_ids=[1, 2, 3])