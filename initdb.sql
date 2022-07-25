create database mkt_data;
\connect mkt_data;
create table spot_info
(
    id                                  serial
        constraint table_name_pk
            primary key,
    language                            varchar,
    region                              varchar,
    "quoteType"                         varchar,
    "typeDisp"                          varchar, 
    "quoteSourceName"                   varchar,
    triggerable                         boolean,
    "customPriceAlertConfidence"        varchar,
    currency                            varchar,
    "shortName"                         varchar,
    "regularMarketChange"               float8,
    "regularMarketChangePercent"        float8,
    "regularMarketPrice"                float8,
    "regularMarketDayHigh"              float8,
    "regularMarketDayLow"               float8,
    "regularMarketPreviousClose"        float8,
    bid                                 float8,
    ask                                 float8,
    "regularMarketOpen"                 float8,
    "fiftyTwoWeekLow"                   float8,
    "fiftyTwoWeekHigh"                  float8,
    "fiftyDayAverage"                   float8,
    "twoHundredDayAverage"              float8,
    "messageBoardId"                    varchar,
    "exchangeTimezoneName"              varchar,
    "exchangeTimezoneShortName"         varchar,
    "gmtOffSetMilliseconds"             bigint,
    market                              varchar,
    "esgPopulated"                      boolean,
    exchange                            varchar,
    "marketState"                       varchar,
    tradeable                           boolean,
    "firstTradeDateMilliseconds"        bigint,
    "priceHint"                         bigint,
    "regularMarketTime"                 bigint,
    "regularMarketDayRange"             varchar,
    "regularMarketVolume"               float8,
    "bidSize"                           float8,
    "askSize"                           float8,
    "fullExchangeName"                  varchar,
    "averageDailyVolume3Month"          float8,
    "averageDailyVolume10Day"           float8,
    "fiftyTwoWeekLowChange"             float8,
    "fiftyTwoWeekLowChangePercent"      float8,
    "fiftyTwoWeekRange"                 varchar,
    "fiftyTwoWeekHighChangePercent"     float8,
    "fiftyDayAverageChange"             float8,
    "fiftyDayAverageChangePercent"      float8,
    "twoHundredDayAverageChange"        float8,
    "twoHundredDayAverageChangePercent" float8,
    "sourceInterval"                    float8,
    "exchangeDataDelayedBy"             float8,
    symbol                              varchar
);

create unique index spot_info_id_uindex
    on spot_info (id);

