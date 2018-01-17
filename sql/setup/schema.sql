DROP SCHEMA IF EXISTS harvest;
CREATE SCHEMA IF NOT EXISTS harvest;
USE harvest;

-- tables
-- Table: app_data
CREATE TABLE app_data (
  attribute VARCHAR(50)  NOT NULL,
  value     VARCHAR(250) NOT NULL,
  CONSTRAINT app_data_pk PRIMARY KEY (attribute)
);

-- Table: client
CREATE TABLE client (
  id         INT          NOT NULL,
  name       VARCHAR(100) NULL,
  is_active  BOOL         NULL,
  address    VARCHAR(500) NULL,
  currency   VARCHAR(100) NULL,
  created_at DATETIME     NOT NULL,
  updated_at DATETIME     NOT NULL,
  CONSTRAINT client_pk PRIMARY KEY (id)
);

-- Table: invoice
CREATE TABLE invoice (
  id              BIGINT         NOT NULL,
  client_id       INT            NOT NULL,
  number          VARCHAR(100)   NULL,
  purchase_order  VARCHAR(100)   NULL,
  amount          DECIMAL(13, 4) NULL,
  due_amount      DECIMAL(13, 4) NULL,
  tax             DECIMAL(13, 4) NULL,
  tax_amount      DECIMAL(13, 4) NULL,
  tax2            DECIMAL(13, 4) NULL,
  tax2_amount     DECIMAL(13, 4) NULL,
  discount        DECIMAL(13, 4) NULL,
  discount_amount DECIMAL(13, 4) NULL,
  subject         VARCHAR(250)   NULL,
  notes           TEXT           NULL,
  currency        VARCHAR(100)   NULL,
  period_start    DATE           NULL,
  period_end      DATE           NULL,
  issue_date      DATE           NULL,
  due_date        DATE           NULL,
  sent_at         DATETIME       NULL,
  paid_at         DATETIME       NULL,
  closed_at       DATETIME       NULL,
  created_at      DATETIME       NOT NULL,
  updated_at      DATETIME       NOT NULL,
  CONSTRAINT invoice_pk PRIMARY KEY (id)
);

-- Table: project
CREATE TABLE project (
  id                                  INT            NOT NULL,
  client_id                           INT            NOT NULL,
  name                                VARCHAR(250)   NULL,
  code                                VARCHAR(50)    NULL,
  is_active                           BOOL           NULL,
  is_billable                         BOOL           NULL,
  is_fixed_fee                        BOOL           NULL,
  bill_by                             VARCHAR(250)   NULL,
  hourly_rate                         DECIMAL(13, 4) NULL,
  budget                              DECIMAL(13, 4) NULL,
  budget_by                           VARCHAR(100)   NULL,
  notify_when_over_budget             BOOL           NULL,
  over_budget_notification_percentage DECIMAL(13, 4) NULL,
  over_budget_notification_date       DATE           NULL,
  show_budget_to_all                  BOOL           NULL,
  cost_budget                         DECIMAL(13, 4) NULL,
  cost_budget_include_expenses        BOOL           NULL,
  fee                                 DECIMAL(13, 4) NULL,
  notes                               TEXT           NULL,
  starts_on                           DATE           NULL,
  ends_on                             DATE           NULL,
  created_at                          DATETIME       NOT NULL,
  updated_at                          DATETIME       NOT NULL,
  CONSTRAINT project_pk PRIMARY KEY (id)
);

-- Table: task
CREATE TABLE task (
  id                  INT            NOT NULL,
  name                VARCHAR(250)   NOT NULL,
  billable_by_default BOOL           NULL,
  default_hourly_rate DECIMAL(13, 4) NULL,
  is_default          BOOL           NULL,
  is_active           BOOL           NULL,
  created_at          DATETIME       NOT NULL,
  updated_at          DATETIME       NOT NULL,
  CONSTRAINT task_pk PRIMARY KEY (id)
);

-- Table: time_entry
CREATE TABLE time_entry (
  id               BIGINT         NOT NULL,
  user_id          INT            NOT NULL,
  client_id        INT            NOT NULL,
  project_id       INT            NOT NULL,
  task_id          INT            NOT NULL,
  invoice_id       BIGINT         NULL,
  spent_date       DATE           NULL,
  hours            DECIMAL(13, 4) NULL,
  notes            TEXT           NULL,
  is_locked        BOOL           NULL,
  locked_reason    VARCHAR(100)   NULL,
  is_closed        BOOL           NULL,
  is_billed        BOOL           NULL,
  timer_started_at DATETIME       NULL,
  started_time     VARCHAR(25)    NULL,
  ended_time       VARCHAR(25)    NULL,
  is_running       BOOL           NULL,
  billable         BOOL           NULL,
  budgeted         BOOL           NULL,
  billable_rate    DECIMAL(13, 4) NULL,
  cost_rate        DECIMAL(13, 4) NULL,
  created_at       DATETIME       NOT NULL,
  updated_at       DATETIME       NOT NULL,
  CONSTRAINT time_entry_pk PRIMARY KEY (id)
);

-- Table: user
CREATE TABLE user (
  id                                INT            NOT NULL,
  first_name                        VARCHAR(100)   NOT NULL,
  last_name                         VARCHAR(100)   NULL,
  email                             VARCHAR(100)   NOT NULL,
  telephone                         VARCHAR(100)   NULL,
  timezone                          VARCHAR(100)   NULL,
  has_access_to_all_future_projects BOOL           NULL,
  is_contractor                     BOOL           NULL,
  is_admin                          BOOL           NULL,
  is_project_manager                BOOL           NULL,
  can_see_rates                     BOOL           NULL,
  can_create_projects               BOOL           NULL,
  can_create_invoices               BOOL           NULL,
  is_active                         BOOL           NULL,
  weekly_capacity                   INT            NULL,
  default_hourly_rate               DECIMAL(13, 4) NULL,
  cost_rate                         DECIMAL(13, 4) NULL,
  avatar_url                        VARCHAR(500)   NULL,
  created_at                        DATETIME       NOT NULL,
  updated_at                        DATETIME       NOT NULL,
  CONSTRAINT user_pk PRIMARY KEY (id)
);

-- foreign keys
-- Reference: invoice_client (table: invoice)
ALTER TABLE invoice
  ADD CONSTRAINT invoice_client FOREIGN KEY invoice_client (client_id)
REFERENCES client (id);

-- Reference: project_client (table: project)
ALTER TABLE project
  ADD CONSTRAINT project_client FOREIGN KEY project_client (client_id)
REFERENCES client (id);

-- Reference: time_entry_client (table: time_entry)
ALTER TABLE time_entry
  ADD CONSTRAINT time_entry_client FOREIGN KEY time_entry_client (client_id)
REFERENCES client (id);

-- Reference: time_entry_invoice (table: time_entry)
ALTER TABLE time_entry
  ADD CONSTRAINT time_entry_invoice FOREIGN KEY time_entry_invoice (invoice_id)
REFERENCES invoice (id);

-- Reference: time_entry_project (table: time_entry)
ALTER TABLE time_entry
  ADD CONSTRAINT time_entry_project FOREIGN KEY time_entry_project (project_id)
REFERENCES project (id);

-- Reference: time_entry_task (table: time_entry)
ALTER TABLE time_entry
  ADD CONSTRAINT time_entry_task FOREIGN KEY time_entry_task (task_id)
REFERENCES task (id);

-- Reference: time_entry_user (table: time_entry)
ALTER TABLE time_entry
  ADD CONSTRAINT time_entry_user FOREIGN KEY time_entry_user (user_id)
REFERENCES user (id);

-- End of file.

