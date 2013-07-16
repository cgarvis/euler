class String
  def palindrome?
    self == self.reverse
  end
end

largest = 0

(100..999).each do |i|
  (100..999).each do |j|
    num = i*j

    if num > largest and num.to_s.palindrome?
      largest = num
    end
  end
end

puts largest
