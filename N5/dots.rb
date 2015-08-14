class ChallengeN5
  def getFromStream(s, res, l)
    s.each do |h|
      d = h[:x] * h[:x] + h[:y] * h[:y] # Find the distance
      h[:d] = d
      if res.size < l
        res << h
      else # replace max coordinates with the smaller one
        res.sort!{|left, right| left[:d] <=> right[:d]}
        if res.last[:d] >= d
          res.pop
          res << h
        end
      end
    end
  end

  def run(stream = [], limit = 0 )
    # stream var mimics a chunk of an endless stream
    res = [] # Stores a result
    getFromStream(stream, res, limit)
    res
  end

end
