#FROM golang:latest AS builder
FROM golang:1.20.2

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

COPY libsqlite3/usr/lib/x86_64-linux-gnu/libsqlite3.so.0.8.6 /usr/lib/x86_64-linux-gnu/
COPY libsqlite3/usr/lib/x86_64-linux-gnu/libsqlite3.so.0 /usr/lib/x86_64-linux-gnu/
RUN ln -s /usr/lib/x86_64-linux-gnu/libsqlite3.so.0 /usr/lib/x86_64-linux-gnu/libsqlite3.so

COPY x64/libcrypto.so.1.0.0 /usr/lib
RUN ln -s /usr/lib/libcrypto.so.1.0.0 /usr/lib/libcrypto.so

COPY x64/libssl.so.1.0.0 /usr/lib
RUN ln -s /usr/lib/libssl.so.1.0.0 /usr/lib/libssl.so

COPY x64/libicudata.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.53 \
    && ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.5 \
    && ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so

COPY x64/libicui18n.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.53 \
    && ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.5 \
    && ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so

COPY x64/libicuuc.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.53 \
    && ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.5 \
    && ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so

COPY x64/libQt5Core.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5.4 \
    && ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5 \
    && ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so

COPY x64/libQt5Network.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5.4 \
    && ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5 \
    && ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so

COPY x64/libQt5Script.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5.4 \
    && ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5 \
    && ln -s /usr/lib/libQt5Script.so


# Add user
RUN useradd -ms /bin/bash ebarimtuser


# Set working directory
WORKDIR /home/ebarimtuser/app


# Copy the rest of the project files into the container at /app
COPY . .

# Download dependencies
RUN go mod download

# Build the project
RUN go build -buildvcs=false -o main .

# Run the application when the container starts
RUN chmod +x /home/ebarimtuser/app/main

# Expose port 3000 for the application
RUN mkdir /home/ebarimtuser/.vatps && \
    chown ebarimtuser:ebarimtuser /home/ebarimtuser/.vatps

RUN chmod -R 755 /home/ebarimtuser/.vatps

RUN chown ebarimtuser:ebarimtuser /home/ebarimtuser/app
USER ebarimtuser
RUN mkdir /home/ebarimtuser/app/solfiles && \
    chown ebarimtuser:ebarimtuser /home/ebarimtuser/app/solfiles
CMD ["./main"]

