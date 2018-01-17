DELIMITER ;

DROP FUNCTION IF EXISTS FN_WORKING_DAYS_IN_MONTH;

DELIMITER ;;

CREATE FUNCTION FN_WORKING_DAYS_IN_MONTH(yyyy_mmm VARCHAR(25))
  RETURNS INT

  BEGIN
    DECLARE StartDate DATE DEFAULT DATE(CONCAT(yyyy_mmm, '-01'));
    DECLARE EndDate DATE DEFAULT DATE_ADD(LAST_DAY(StartDate), INTERVAL 1 DAY);

    RETURN (
      5 * (DATEDIFF(EndDate, StartDate) DIV 7) +
      MID('0123455501234445012333450122234501101234000123450', 7 * WEEKDAY(StartDate) + WEEKDAY(EndDate) + 1, 1)
    );
  END ;;

DELIMITER ;