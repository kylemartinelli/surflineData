
DROP TABLE IF EXISTS spots CASCADE;

CREATE TABLE IF NOT EXISTS spots(
  id  SERIAL PRIMARY KEY,
  surflineId VARCHAR(255) NOT NULL UNIQUE,
  surflineUrl VARCHAR(255) NOT NULL,
  spotName VARCHAR(255) NOT NULL,
  lat DECIMAL(8,6),
  lon DECIMAL(9,6),
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS idealConditions;

CREATE TABLE IF NOT EXISTS idealConditions(
  id SERIAL PRIMARY KEY,
  spotId INT NOT NULL,
  idealSeason VARCHAR(255)[],
  idealSeasonDescription VARCHAR(255) NOT NULL,
  idealTide VARCHAR(255)[],
  idealTideDescription VARCHAR(255) NOT NULL,
  idealSize VARCHAR(255)[],
  idealSizeDescription VARCHAR(255) NOT NULL,
  idealWindDirections VARCHAR(255)[],
  idealWindDirectionsDescription VARCHAR(255) NOT NULL,
  idealSwellDirections VARCHAR(255)[],
  idealSwellDirectionsDescription VARCHAR(255) NOT NULL,


  FOREIGN KEY(spotId) REFERENCES spots(id) ON DELETE CASCADE
);


DROP TABLE IF EXISTS spotCharacteristics;
CREATE TABLE IF NOT EXISTS spotCharacteristics(
  id SERIAL PRIMARY KEY,
  spotId INT NOT NULL,
  crowdFactorDescription VARCHAR(255) NOT NULL,
  crowdFactorRating INT NOT NULL,
  crowdFactorSummary VARCHAR(255) NOT NULL,

  localVibeDescription VARCHAR(255) NOT NULL,
  localVibeRating INT NOT NULL,
  localVibeSummary VARCHAR(255) NOT NULL,

  shoulderBurnDescription VARCHAR(255) NOT NULL,
  shoulderBurnRating INT NOT NULL,
  shoulderBurnSummary VARCHAR(255) NOT NULL,

  spotRatingDescription VARCHAR(255) NOT NULL,
  spotRatingRating INT NOT NULL,
  spotRatingSummary VARCHAR(255) NOT NULL,

  waterQualityDescription VARCHAR(255) NOT NULL,
  waterQualityRating INT NOT NULL,
  waterQualitySummary VARCHAR(255) NOT NULL,

  abilityLevels VARCHAR(255)[],
  abilityDescription VARCHAR(255) NOT NULL,

  bottomType VARCHAR(255)[],
  bottomTypeDescription VARCHAR(255) NOT NULL,

  boardTypes  VARCHAR(255)[],

  breakType  VARCHAR(255)[],
  spotDescription TEXT,

  FOREIGN KEY (spotId) REFERENCES spots(id)

);


DROP TABLE IF EXISTS spotBreadCrumbs;

CREATE TABLE IF NOT EXISTS spotBreadCrumbs(
  id SERIAL PRIMARY KEY,
  spotId INT NOT NULL,
  country VARCHAR(255) NOT NULL,
  countryId INT NOT NULL,
  region VARCHAR(255) NOT NULL,
  regionId INT NOT NULL,
  subRegion VARCHAR(255) NOT NULL,
  subRegionId INT NOT NULL,
  area VARCHAR(255) DEFAULT NULL,
  areaId INT DEFAULT NULL,


  FOREIGN KEY(spotId) REFERENCES spots(id)
);


-- mysql -u root -h localhost -p < db/surflineData.sql
-- psql -U kylemartinelli surflinedata -a -f db/surflineData.sql