import grpc
from concurrent import futures
import threading
from kafka import KafkaConsumer
import json
import os

from .recommendation_pb2_grpc import add_RecommendationServiceServicer_to_server
from .servicers.recommendation import RecommendationService


# Kafka configuration
KAFKA_URL = os.getenv('KAFKA_URL')
KAFKA_USERNAME = os.getenv('KAFKA_USERNAME')
KAFKA_PASSWORD = os.getenv('KAFKA_PASSWORD')
KAFKA_ARTICLE_TOPIC = os.getenv('KAFKA_ARTICLE_INTERACTION_TOPIC', 'article-interactions')
KAFKA_PRODUCT_TOPIC = os.getenv('KAFKA_PRODUCT_INTERACTION_TOPIC', 'product-interactions')

def process_article_interaction(interaction):
    # TODO: Implement logic to update article-product associations
    print(f"Processing article interaction: {interaction}")

def process_product_interaction(interaction):
    # TODO: Implement logic to update user-product preferences
    print(f"Processing product interaction: {interaction}")

def kafka_consumer_worker():
    consumer = KafkaConsumer(
        KAFKA_ARTICLE_TOPIC,
        KAFKA_PRODUCT_TOPIC,
        bootstrap_servers=[KAFKA_URL],
        sasl_mechanism='SCRAM-SHA-256',
        security_protocol='SASL_SSL',
        sasl_plain_username=KAFKA_USERNAME,
        sasl_plain_password=KAFKA_PASSWORD,
        group_id='recommendation-processor',
        auto_offset_reset='earliest',
        value_deserializer=lambda x: json.loads(x.decode('utf-8'))
    )

    print("Kafka consumer started. Waiting for messages...")

    try:
        for message in consumer:
            topic = message.topic
            interaction = message.value

            if topic == KAFKA_ARTICLE_TOPIC:
                process_article_interaction(interaction)
            elif topic == KAFKA_PRODUCT_TOPIC:
                process_product_interaction(interaction)
            else:
                print(f"Unknown topic: {topic}")
    except KeyboardInterrupt:
        pass
    finally:
        consumer.close()

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_RecommendationServiceServicer_to_server(RecommendationService(), server)
    port = 50051
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    print(f"gRPC Server started on port {port}")

    # Start Kafka consumer in a separate thread
    kafka_thread = threading.Thread(target=kafka_consumer_worker, daemon=True)
    kafka_thread.start()

    server.wait_for_termination()

def main():
    serve()

if __name__ == '__main__':
    main()