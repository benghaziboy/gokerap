package rapper

const (
    createRapperTable = `
    CREATE TABLE rappers (
        id      SERIAL PRIMARY KEY,
        name    VARCHAR UNIQUE NOT NULL
    );
    `
    
    createRapper = `
    INSERT INTO rappers (name) VALUES ($1) RETURNING id;
    `
    
    findRapper = `
    SELECT id, name
    FROM rappers
    WHERE name = $1;
    `

    getRapper = `
    SELECT id, name
    FROM rappers
    WHERE id = $1
    `

    listRapper = `
    SELECT id, name
    FROM rappers
    `

    updateRapper = `
    UPDATE rappers name = $1 WHERE id = $2;
    `
)
