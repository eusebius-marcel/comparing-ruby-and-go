class Hamming
  def self.distance(strand1, strand2)
    if strand1.length != strand2.length
      raise ArgumentError, "Both strands don't have the same length"
    end

    strand1.length.times.count do |index|
      strand1[index] != strand2[index]
    end
  end
end
