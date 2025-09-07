param(
    [int]$count = 10000,
    [string]$url = 'http://localhost:64333/get',
    [int]$progressStep = 1000,
    [int]$timeoutSec = 5,
    [int]$retries = 2
)

$sum0 = [double]$count
$sum1 = 0.0
$failed = 0

Write-Host "=== Checking RES ==="
Write-Host "count = $count, url = $url"
Write-Host "SUM0 (number of elements) = $sum0"
Write-Host ""

$sw = [System.Diagnostics.Stopwatch]::StartNew()

for ($i = 1; $i -le $count; $i++) {
    $x = 1.0
    $attempt = 0
    $got = $false
    $multiplier = 0.0

    while (-not $got -and $attempt -le $retries) {
        try {
            $attempt++
            $resp = Invoke-RestMethod -Uri $url -Method GET -TimeoutSec $timeoutSec
            $multiplier = [double]$resp.result
            $got = $true
        }
        catch {
            if ($attempt -gt $retries) {
                Write-Host "[$i] ERROR: failed after $attempt attempts." -ForegroundColor Yellow
                $failed++
            }
            else { Start-Sleep -Milliseconds 100 }
        }
    }

    if ($got) {
        if ($multiplier -gt $x) {
            $sum1 += $x
        }
    }

    if (($i % $progressStep) -eq 0) {
        $pct = [math]::Round($i / $count * 100, 2)
        Write-Host ("Progress: {0}/{1} ({2}%)  elapsed: {3}  failed: {4}" -f $i, $count, $pct, $sw.Elapsed.ToString(), $failed)
    }
}

$sw.Stop()
$RTP = 0.0
if ($sum0 -gt 0) { $RTP = $sum1 / $sum0 }

Write-Host ""
Write-Host "-------------------------------------------"
Write-Host "Requested: $count"
Write-Host "Failed requests: $failed"
Write-Host ("SUM0 (number of elements) = {0:N0}" -f $sum0)
Write-Host ("SUM1 (after filtering) = {0:N6}" -f $sum1)
Write-Host ("RES = {0:N6}" -f $RTP)
Write-Host ("Elapsed time: {0}" -f $sw.Elapsed)
Write-Host "-------------------------------------------"
