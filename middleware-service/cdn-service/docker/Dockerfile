FROM golang:1.14
# Name service
LABEL Name=middleware-cdn-service

ARG LIBVIPS_VERSION=8.9.2
ARG GOLANGCILINT_VERSION=1.23.3

# Installs libvips + required libraries
RUN DEBIAN_FRONTEND=noninteractive \
  apt-get update && \
  apt-get install --no-install-recommends -y \
  ca-certificates \
  automake build-essential curl \
  gobject-introspection gtk-doc-tools libglib2.0-dev libjpeg62-turbo-dev libpng-dev \
  libwebp-dev libtiff5-dev libgif-dev libexif-dev libxml2-dev libpoppler-glib-dev \
  swig libmagickwand-dev libpango1.0-dev libmatio-dev libopenslide-dev libcfitsio-dev \
  libgsf-1-dev fftw3-dev liborc-0.4-dev librsvg2-dev libimagequant-dev libheif-dev && \
  cd /tmp && \
  curl -fsSLO https://github.com/libvips/libvips/releases/download/v${LIBVIPS_VERSION}/vips-${LIBVIPS_VERSION}.tar.gz && \
  tar zvxf vips-${LIBVIPS_VERSION}.tar.gz && \
  cd /tmp/vips-${LIBVIPS_VERSION} && \
	CFLAGS="-g -O3" CXXFLAGS="-D_GLIBCXX_USE_CXX11_ABI=0 -g -O3" \
    ./configure \
    --disable-debug \
    --disable-dependency-tracking \
    --disable-introspection \
    --disable-static \
    --enable-gtk-doc-html=no \
    --enable-gtk-doc=no \
    --enable-pyvips8=no \
    --prefix=/vips && \
  make && \
  make install && \
  ldconfig
RUN go get -u golang.org/x/lint/golint
# Enable gomod & other env
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# Setup working directory /build
WORKDIR /build
# Copy source code into the container
ADD . /build
# Copy & download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
# Build the service
RUN go build -o main .
# Move to /dist directory - place for resulting binaries
WORKDIR /dist
# Copy binary from build to main folder
RUN cp /build/main .
# Export necessary port
EXPOSE 8080
# Command to run service container when the container starts
CMD ["/dist/main"]


# Service prod image
FROM alpine:3.9.2 as prod
# Setup working directory /root/
WORKDIR /root/
# Copy service resulting binaries
COPY --from=dev /dist .
# Service listens on port 8080.
EXPOSE 8080
# Run product service container when the container starts
ENTRYPOINT ["./cdn"]



