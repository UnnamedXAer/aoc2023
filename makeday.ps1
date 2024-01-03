param([int] $dayNo = -1)

if ($dayNo -lt 1 ) {
  Write-Host "missing day number"
  exit 1
}

# mkdir "./day$dayNo"
	
New-Item -Path  "./" -Name "day$dayNo" -ItemType Directory


if (!$?) {
  Write-Host "creating ./day$dayNo failed with code $LASTEXITCODE"
  exit 1
}



"" >> "./day$dayNo/data.txt"
"" >> "./day$dayNo/data_t.txt"

@"
package day${dayNo}

func Ex1() {

}

func Ex2() {

}
"@ >> "./day$dayNo/day$dayNo.go" 

Write-Host "Day $dayNo created."

code "./day$dayNo/day$dayNo.go"
code "./day$dayNo/data.txt"
code "./day$dayNo/data_t.txt"