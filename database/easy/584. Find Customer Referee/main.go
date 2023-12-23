SET ANSI_NULLS OFF;
SELECT name from Customer WHERE referee_id != 2 OR referee_id = NULL;
