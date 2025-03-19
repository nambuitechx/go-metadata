FROM python:3.9-slim

# Install curl (or any other necessary packages)
RUN apt-get update && apt-get install -y curl

RUN rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy requirements.txt and install dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application code
COPY . .

# Command to run your FastAPI application
CMD ["python","main.py"]
