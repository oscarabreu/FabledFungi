FROM amazonlinux:2

RUN yum install -y python3 zip
RUN python3 -m pip install --upgrade pip

# Set up a working directory to hold the application code within the image
WORKDIR /var/task

# Copy the function code and requirements file into the image
COPY lambda_function.py .
COPY requirements.txt .

# Install the Python dependencies from requirements.txt to the Docker image
RUN pip install -r requirements.txt -t .

# Package everything into a zip file
RUN zip -r9 /var/task/function.zip .
