#!/usr/bin/env python3

import os
import subprocess
import sys
import argparse

def check_docker_compose_file(path="."):
    yml_file = os.path.join(path, "docker-compose.yml")
    yaml_file = os.path.join(path, "docker-compose.yaml")
    
    if os.path.exists(yml_file) or os.path.exists(yaml_file):
        pass
    else:
        print(f"Error: Missing docker-compose.yml or docker-compose.yaml file in the path: {os.path.abspath(path)}")
        sys.exit(1)

def run_docker_compose(profile):
    service_name = "auth-service-" + profile
    try:
        result = subprocess.run(
            ["docker", "compose", "--profile", profile, "up", "--build", "--exit-code-from", service_name],
            check=False
        )
        return result.returncode
    except subprocess.CalledProcessError as e:
        print(f"Error: {e}")
        sys.exit(1)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Run docker-compose with a specified profile.")
    parser.add_argument(
        "profile",
        nargs="?",
        default="dev",
        help="Docker Compose profile to use (default: 'dev')"
    )
    args = parser.parse_args()

    if args.profile == "dev":
        print("Warning: Using 'dev' profile as default.")

    check_docker_compose_file()

    exit_code = run_docker_compose(args.profile)
    sys.exit(exit_code)

