require_relative '../dots'

describe "Dots challenge " do
  before do
    @c = ChallengeN5.new
  end

  context "validating logic" do
    it "should generate right result" do
      res = @c.run([{:x => -5, :y => 5}, {:x => 5, :y => 5}, {:x => 1, :y => -1}, {:x => 10, :y => 5}], 2)
      expect(res).to match_array([{:x=>1, :y=>-1, :d=>2}, {:x=>-5, :y=>5, :d=>50} ])
    end
    it "should generate right result" do
      res = @c.run([{:x => -5, :y => 5}, {:x => 5, :y => 5}, {:x => 1, :y => -1}, {:x => 0, :y => 1}], 2)
      expect(res).to match_array([{:x=>1, :y=>-1, :d=>2}, {:x=>0, :y=>1, :d=>1} ])
    end
  end
end
