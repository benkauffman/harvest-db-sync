USE harvest;

SET @StartDate := '2017-01-01';
SET @EndDate := '2017-12-31';

SET @PartTimeEmployees := 'Example';
SET @QueryLastNames := 'Example';

SELECT
  *,
  total_hrs / month_total_working_hours AS percent_working_effectiveness
FROM (
       SELECT
         *,
         total_days_in_month * (CASE WHEN last_name IN (@PartTimeEmployees)
           THEN 4
                                ELSE 8 END) AS month_total_working_hours

       FROM (

              SELECT
                *,
                billable_hrs / total_hrs                 AS percent_billable_effectiveness,
                FN_WORKING_DAYS_IN_MONTH(year_and_month) AS total_days_in_month
              FROM (
                     SELECT
                       DATE_FORMAT(spent_date, '%Y-%m') AS year_and_month,
                       first_name,
                       last_name,
                       SUM(CASE WHEN billable IS TRUE
                         THEN hours
                           ELSE 0 END)                  AS billable_hrs,
                       SUM(CASE WHEN billable IS FALSE
                         THEN hours
                           ELSE 0 END)                  AS non_billable_hrs,
                       SUM(hours)                       AS total_hrs
                     FROM (SELECT *
                           FROM user
                           WHERE
                             last_name IN
                             (@QueryLastNames)) AS u
                       LEFT JOIN time_entry AS te ON u.id = te.user_id

                     WHERE te.spent_date >= @StartDate AND te.spent_date <= @EndDate
                     GROUP BY first_name, last_name, year_and_month
                   ) AS smry

            ) AS mnth_smry
     ) AS final




