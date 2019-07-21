/*
In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

Example:

HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
Notes:

All numbers are valid Int32, no need to validate them.
There will always be at least one number in the input string.
Output string must be two numbers separated by a single space, and highest number is first.
*/

public static String HighAndLow(String numbers) {
		
        Integer max = 0;
        Integer min = 0;
        
        if (numbers.length() == 0)
        	return "";

        String[] nums = numbers.split(" ");
        
        for (String num : nums)
        {
            if (Integer.valueOf(num) >= max)
            	max = Integer.valueOf(num);
            
            if (Integer.valueOf(num) <= min)
            	min = Integer.valueOf(num);
        }

        return "" + max + " " + min;
        
	}