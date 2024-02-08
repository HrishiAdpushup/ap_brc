#!/usr/bin/bash
#start=echo `date`;
current_date=$(date +"%Y-%m-%d %H:%M:%S")

awk -F';' '{
    cityCount[$1]++;
    sumOfTempForEachCity[$1] += $2;
    totalTemperatur += $2
}
END {
    # Print the frequency of each substring
    for (city in cityCount) {
         totalCount++
        #print "freq of " city ": " cityCount[city]
        print "Sum of " city ": " sumOfTempForEachCity[city];
        
        print "average of " city ": " sumOfTempForEachCity[city]/cityCount[city]
    }
        print "total is " totalCount 
        print "Average Temp is " totalTemperatur/totalCount
}' measurements.txt

echo $current_date
echo `date`

