actual = str(input()).split(' ')

a_day = int(actual[0])
a_month = int(actual[1])
a_year = int(actual[2])

exepcted = str(input()).split(' ')

e_day = int(exepcted[0])
e_month = int(exepcted[1])
e_year = int(exepcted[2])


fine = 0


if a_year > e_year:
    fine = 10000

if a_year == e_year:

    if a_month == e_month and e_day < a_day:
        fine = 15 * (a_day - e_day)

    if a_month > e_month:
        fine = 500 * (a_month - e_month)


print(fine)