# Base aspnet core image
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-buster-slim AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

# Build step container image
FROM mcr.microsoft.com/dotnet/core/sdk:3.1-buster AS build
# Setup working directory /build
WORKDIR /src
# Copy projects
COPY web-app-service.sln ./
COPY Core/*.csproj ./Core/
COPY Infrastructure/*.csproj ./Infrastructure/
COPY Web/*.csproj ./Web/
# Run projects restore cmd
RUN dotnet restore
COPY . .
# Build projects
WORKDIR "/src/Core"
RUN dotnet build -c Release -o /app
WORKDIR "/src/Infrastructure"
RUN dotnet build -c Release -o /app
WORKDIR "/src/Web"
RUN dotnet build -c Release -o /app

# Publish step container image
FROM build AS publish
RUN dotnet publish -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
EXPOSE 88
ENTRYPOINT ["dotnet", "Web.dll"]