$containerName = "gocadastroproduto-postgres-1"

# Check if the Docker container is running
$containerRunning = docker ps -q -f "name=$containerName" -f "status=running"

if ($containerRunning) {
    Write-Host "Container $containerName is running."
} else {
    # Start the Docker containers
    docker-compose up -d

    # Function to wait for the PostgreSQL container to be ready
    function WaitForPostgres {
        $maxAttempts = 30  # Maximum number of attempts to check if PostgreSQL is ready
        $attempts = 0

        # Loop to check if PostgreSQL is ready
        while ($attempts -lt $maxAttempts) {
            $result = docker inspect --format="{{json .State.Health.Status}}" $containerName
            if ($result -eq '"healthy"') {
                Write-Host "PostgreSQL is ready!"
                return
            }

            Write-Host "Waiting for PostgreSQL to start..."
            $attempts++
            Start-Sleep -Seconds 2
        }

        Write-Host "Timeout: PostgreSQL did not become ready within the specified time."
    }

    # Wait for PostgreSQL to be available
    WaitForPostgres

    # Execute a bash command in the PostgreSQL container
    docker exec -it $containerName bash -c "sh migration/ExecuteInsert.sh"
}




