CREATE TYPE origin AS ENUM ('top', 'left');

CREATE TYPE request_status as ENUM ('init', 'retry', 'success', 'error');

CREATE TABLE game (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    ended_at TIMESTAMP WITH TIME ZONE,
    requests INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE request (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'), 
    ended_at TIMESTAMP WITH TIME ZONE,
    game_id uuid REFERENCES game(id),
    colour CHAR(7) NOT NULL,
    origin origin NOT NULL,
    speed INTEGER NOT NULL CHECK (speed >= 1 AND speed <= 10),
    width INTEGER NOT NULL CHECK (width >= 1 AND speed <= 10),
    status request_status NOT NULL
);
