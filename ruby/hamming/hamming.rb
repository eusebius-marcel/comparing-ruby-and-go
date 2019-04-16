class Hamming
  def self.compute(strand1, strand2)
    raise ArgumentError, "Both strands don't have the same length" if strand1.length != strand2.length

    strand1.length.times.count do |nucleotide|
      strand1[nucleotide] != strand2[nucleotide]
    end
  end
end
