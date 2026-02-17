CREATE TABLE IF NOT EXISTS customers (
  id VARCHAR(20) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100),
  phone VARCHAR(30),
  credit_limit NUMERIC(15,2) DEFAULT 0,
  credit_used NUMERIC(15,2) DEFAULT 0,
  currency VARCHAR(5) DEFAULT 'USD',
  status VARCHAR(20) DEFAULT 'ACTIVE',
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS orders (
  id VARCHAR(20) PRIMARY KEY,
  customer_id VARCHAR(20) REFERENCES customers(id),
  total_amount NUMERIC(15,2),
  currency VARCHAR(5) DEFAULT 'USD',
  status VARCHAR(20) DEFAULT 'PENDING',
  items JSONB,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS materials (
  id VARCHAR(20) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  uom VARCHAR(20),
  price NUMERIC(15,2),
  currency VARCHAR(5) DEFAULT 'USD',
  stock INTEGER DEFAULT 0,
  plant VARCHAR(10),
  created_at TIMESTAMP DEFAULT NOW()
);

-- Seed customers
INSERT INTO customers (id, name, email, phone, credit_limit, credit_used, currency)
VALUES
  ('C-001', 'PT Maju Bersama', 'finance@majubersama.co.id', '+6221-555-1001', 100000.00, 35000.00, 'USD'),
  ('C-002', 'CV Teknologi Nusantara', 'ap@teknusa.co.id', '+6221-555-1002', 50000.00, 50000.00, 'USD'),
  ('C-003', 'PT Sumber Rezeki', 'procurement@sumberrezeki.com', '+6221-555-1003', 75000.00, 10000.00, 'USD')
ON CONFLICT DO NOTHING;

-- Seed materials
INSERT INTO materials (id, name, description, uom, price, currency, stock, plant)
VALUES
  ('MAT-001', 'Steel Rod 10mm', 'Carbon steel rod diameter 10mm', 'KG', 2.50, 'USD', 5000, 'PLANT-JKT'),
  ('MAT-002', 'Copper Wire 2.5mm', 'Electrical copper wire 2.5mm', 'MTR', 1.80, 'USD', 12000, 'PLANT-JKT'),
  ('MAT-003', 'Industrial Paint 5L', 'Epoxy industrial paint 5L', 'CAN', 45.00, 'USD', 800, 'PLANT-SBY')
ON CONFLICT DO NOTHING;