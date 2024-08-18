import sys
import os
from grpc_tools import protoc
import subprocess

def generate_proto():
    # Get the current directory
    current_dir = os.path.dirname(os.path.abspath(__file__))
    parent_dir = os.path.dirname(current_dir)
    
    protoc.main([
        'grpc_tools.protoc',
        f'-I{parent_dir}/protos',
        f'--python_out={current_dir}',
        f'--pyi_out={current_dir}',
        f'--grpc_python_out={current_dir}',
        f'{parent_dir}/protos/recommendation.proto'
    ])
    
    # Fix imports in the generated files
    pb2_grpc_file = os.path.join(current_dir, 'recommendation_pb2_grpc.py')
    with open(pb2_grpc_file, 'r') as f:
        content = f.read()
    
    content = content.replace(
        'import recommendation_pb2 as recommendation__pb2',
        'from . import recommendation_pb2 as recommendation__pb2'
    )
    
    with open(pb2_grpc_file, 'w') as f:
        f.write(content)
    
    print(f"Files generated in: {current_dir}")
    print(f"Contents of {current_dir}:")
    print(os.listdir(current_dir))


def run_dev_server():
    generate_proto()  # Generate proto files before starting the server
    subprocess.run([
        "python", "-m", "recommendation.server"
    ])

if __name__ == "__main__":
    sys.exit(generate_proto())