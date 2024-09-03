
DROP TABLE IF EXISTS spots CASCADE;

CREATE TABLE IF NOT EXISTS spots(
  id  SERIAL PRIMARY KEY,
  surflineId VARCHAR(2000) NOT NULL UNIQUE,
  surflineUrl VARCHAR(2000) NOT NULL,
  spotName VARCHAR(2000) NOT NULL,
  lat DECIMAL(8,6),
  lon DECIMAL(9,6),
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS idealConditions;

CREATE TABLE IF NOT EXISTS idealConditions(
  id SERIAL PRIMARY KEY,
  spotId INT NOT NULL,
  idealSeason VARCHAR(2000)[],
  idealSeasonDescription VARCHAR(2000) NOT NULL,
  idealTide VARCHAR(2000)[],
  idealTideDescription VARCHAR(2000) NOT NULL,
  idealSize VARCHAR(2000)[],
  idealSizeDescription VARCHAR(2000) NOT NULL,
  idealWindDirections VARCHAR(2000)[],
  idealWindDirectionsDescription VARCHAR(2000) NOT NULL,
  idealSwellDirections VARCHAR(2000)[],
  idealSwellDirectionsDescription VARCHAR(2000) NOT NULL,


  FOREIGN KEY(spotId) REFERENCES spots(id) ON DELETE CASCADE
);

ALTER TABLE idealconditions ALTER COLUMN idealseasondescription type text;
ALTER TABLE idealconditions ALTER COLUMN idealtidedescription type text;
ALTER TABLE idealconditions ALTER COLUMN idealwinddirectionsdescription type text;
ALTER TABLE idealconditions ALTER COLUMN idealsizedescription type text;


DROP TABLE IF EXISTS spotCharacteristics;
CREATE TABLE IF NOT EXISTS spotCharacteristics(
  id SERIAL PRIMARY KEY,
  spotId INT NOT NULL,
  crowdFactorDescription TEXT NOT NULL,
  crowdFactorRating INT NOT NULL,
  crowdFactorSummary TEXT NOT NULL,

  localVibeDescription TEXT NOT NULL,
  localVibeRating INT NOT NULL,
  localVibeSummary TEXT NOT NULL,

  shoulderBurnDescription TEXT NOT NULL,
  shoulderBurnRating INT NOT NULL,
  shoulderBurnSummary TEXT NOT NULL,

  spotRatingDescription TEXT NOT NULL,
  spotRatingRating INT NOT NULL,
  spotRatingSummary TEXT NOT NULL,

  waterQualityDescription TEXT NOT NULL,
  waterQualityRating INT NOT NULL,
  waterQualitySummary TEXT NOT NULL,

  abilityLevels TEXT [],
  abilityDescription TEXT NOT NULL,

  bottomType TEXT[],
  bottomTypeDescription TEXT NOT NULL,

  boardTypes  VARCHAR(2000)[],

  breakType  VARCHAR(2000)[],
  spotDescription TEXT,

  hazards TEXT,

  FOREIGN KEY (spotId) REFERENCES spots(id)

);


DROP TABLE IF EXISTS spotBreadCrumbs;

CREATE TABLE IF NOT EXISTS spotBreadCrumbs(
  id SERIAL PRIMARY KEY,
  spotId INT NOT NULL,
  country VARCHAR(255) NOT NULL,
  countryUrl VARCHAR(255) NOT NULL,
  region VARCHAR(255) NOT NULL,
  regionUrl VARCHAR(255) NOT NULL,
  subRegion VARCHAR(255) NOT NULL,
  subRegionUrl VARCHAR(255) NOT NULL,
  area VARCHAR(255) DEFAULT NULL,
  areaUrl VARCHAR(255) DEFAULT NULL,


  FOREIGN KEY(spotId) REFERENCES spots(id)
);


-- mysql -u root -h localhost -p < db/surflineData.sql
-- psql -U kylemartinelli surflinedata -a -f db/surflineData.sql