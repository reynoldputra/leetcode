SELECT emu.unique_id, em.name FROM Employees em 
LEFT JOIN EmployeeUNI emu ON emu.id = em.id;
