const readline = require("readline");

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

class Device {
  constructor() {
    this.x = 0;
    this.y = 0;
  }
}

class Distance {
  constructor() {
    this.x = 0;
    this.y = 0;
  }
}

PrintLinkStationsInfo();

let device = new Device();
let target = { x: 0, y: 0, power: 0 };

console.log(
  "\n\nhint, values should not be in string format, only numbers are accepted"
);

rl.question("Enter X value for device: ", (xInput) => {
  device.x = parseInt(xInput);

  if (!isNaN(device.x)) {
    rl.question("Enter Y value for device: ", (yInput) => {
      device.y = parseInt(yInput);

      if (!isNaN(device.y)) {
        let stations = GetLinkStations();

        // Go through all link-stations.
        stations.forEach((station) => {
          // calculate distance between device and station.
          var distance = CalculateDistance(device, station);

          // make the power zero when distance bigger than reach.
          if (distance > station.reach) {
            if (target.power === 0) {
              target.power = 0;
            }
          } else {
            // read new power value
            var nextPower = CalucatePower(station.reach, distance);

            // compare current power with new one.
            // getting the max power for certain link station and make it (most power)
            if (target.power <= nextPower) {
              target.power = nextPower;
              target.x = station.x;
              target.y = station.y;
            }
          }
        });

        // show a message contain power, most power link-stationin for device point in console.
        if (target.power !== 0) {
          // just for coloring purpose using info
          console.info(
            `\nBest link station for point x:${device.x}, y:${device.y} is x:${target.x}, y:${target.y} with power ${target.power}\n`
          );

          // Not found a link station
        } else if (target.power === 0) {
          console.log(
            `\nNo link station within reach for point x:${device.x}, y:${device.y}`
          );
        }

        rl.close();
      } else {
        console.error("\n\nExpceted integer. The Y value for device");
        rl.close();
      }
    });
  } else {
    console.error("\n\nExpceted integer. The X value for device");
    rl.close();
  }
});

/*
#################################################################
#																																#
#													Utils																	#
#																																#
#################################################################	
*/

// Generate a list of link station which is consist of the following X, Y, and Reach values
function GetLinkStations() {
  let linkStation = [
    { x: 0, y: 0, reach: 3 },
    { x: 20, y: 20, reach: 5 },
    { x: 10, y: 0, reach: 12 },
    { x: 17, y: 17, reach: 2 },
    { x: 30, y: 34, reach: 6 },
    { x: 41, y: 43, reach: 4 },
    { x: 1, y: 3, reach: 4 },
  ];
  return linkStation;
}

// Print the list of link station information.
function PrintLinkStationsInfo() {
  console.log("List of link station:");

  let linkStations = GetLinkStations();
  linkStations.forEach((station) => {
    console.log(
      `X: ${station.x},\t\tY: ${station.y},\t\tReach: ${station.reach}`
    );
  });
}

/*
#################################################################
#																																#
#													Calculations													#
#																																#
#################################################################	
*/

// calculate how far a device from link station
function CalculateDistance(device, station) {
  var distance = new Distance();

  distance.x = station.x - device.x;
  distance.y = station.y - device.y;

  return Math.floor(
    Math.sqrt(Math.pow(distance.x, 2) + Math.pow(distance.y, 2))
  );
}

// Calculate the power, by give a reach of link station and distance as arguments.
function CalucatePower(reach, distance) {
  return Math.pow(reach - distance, 2);
}
