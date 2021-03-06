from math import ceil, log
def nb_year(p0, percent, aug, p):
    if not percent: return ceil(1.*(p - p0) / aug)
    percent = 1 + percent / 100.
    r = aug / (1 - percent)
    return ceil(log((p - r) / (p0 - r), percent))


p0 = 1500
percent = 5
aug = 100
p = 5000

print(nb_year(p0,percent,aug,p))