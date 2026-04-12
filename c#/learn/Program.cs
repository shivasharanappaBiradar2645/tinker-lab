using System;

namespace Learn
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello, World shiva");

            int age = 20;
            Console.WriteLine(age);
            long bigNumber = 900000L;
            // like java L long, D double(default) , F float
            // decimal M

            string name = "shiva";
            char firstC = 's';

            string textAge = "20";
            age = Convert.ToInt16(textAge);

            const bool isMale = true;

            Console.Write("Enter college name: ");
            string collegeName = Console.ReadLine();


            Console.WriteLine("I am " + name + " from " + collegeName);
            Console.WriteLine("My age is " + Convert.ToString(age));

            if (age >= 18)
            {
                Console.WriteLine("I can vote");
            }
            else
            {
                Console.WriteLine("I cannot vote");
            }

            Console.Write("Enter the day of the week: ");
            int day = Convert.ToInt16(Console.ReadLine());
            switch (day)
            {
                case 1:
                    Console.WriteLine("Monday");
                    break;

                case 2:
                    Console.WriteLine("Tuesday");
                    break;

                default:
                    Console.WriteLine("default");
                    break;
            }

            for (int i = 0; i < 5; i++)
            {
                Console.WriteLine("hip hip hooray " + Convert.ToString(i));
            }

            while (age < 100)
            {
                age++;
            }

            // condition ? true : false;

            double value = age / 3D;
            Console.WriteLine(string.Format("the value is {0:.00000}",value));//the zero important like format in python
            Console.WriteLine("The value is {0} ", value);

            if (!int.TryParse(textAge, out age)){
                Console.Write("failed");
            }

            string path = @"~/tinker-lab/c#/learn"; //like raw string in python for paths but called verbatism

	//used .equals for string not ==
	
	   int[] numb = new int[3];
	  numb[0]=56;
	 foreach(int num in numb)
	 { 
		
		


        }
	 //Array.Sort(numb)
	
	List<int> listnum = new List<int>();
       listnum.Add(52);
	
	Dictionary<int,string> names = new Dictionary<int,string>();
 	

	}
}
