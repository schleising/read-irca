FROM python:3.9-slim-buster
ENV PYTHONUNBUFFERED 1

# Install the build tools
RUN apt update && apt install -y build-essential libgdal-dev git

# Make the code directory
# RUN mkdir /workspaces
# WORKDIR /code

# Install requirements for the covid charts script
# COPY requirements.txt /code/
# RUN pip install -r requirements.txt

# Run a command to ensure the container does not exit
# ENTRYPOINT [ "tail" ]
# CMD [ "-f", "/dev/null" ]
