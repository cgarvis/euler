# The prime factors of 13195 are 5, 7, 13 and 29.
#
# What is the largest prime factor of the number 600851475143 ?

n = 600851475143

(1..n).each do |factor|
  if n % factor == 0
    n /= factor
  end

  if n == 1
    puts factor
    break
  end
end
