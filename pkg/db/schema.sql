-- Orders table
CREATE TABLE orders (
                        id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                        name VARCHAR NOT NULL);
