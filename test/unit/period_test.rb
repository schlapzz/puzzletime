require File.dirname(__FILE__) + '/../test_helper'

class PeriodTest < Test::Unit::TestCase
  :masterdatas

  def setup
    @half_year = Period.new(Date.new(2006,1,1), Date.new(2006,6,30))
    @one_month = Period.new(Date.new(2006,3,1), Date.new(2006,3,31))
    @two_month = Period.new(Date.new(2005,12,1), Date.new(2006,1,31))
    @one_day = Period.new(Date.new(2006,1,3), Date.new(2006,1,3))
    @holy_day = Period.new(Date.new(2006,1,1), Date.new(2006,1,1))
  end

  # Replace this with your real tests.
  def test_musttime
    assert_equal @half_year.musttime, 129 * 8
    assert_equal @one_month.musttime, 23 * 8
    assert_equal @two_month.musttime, 42 * 8
    assert_equal @one_day.musttime, 8
    assert_equal @holy_day.musttime, 0
  end
  
  def test_length
    assert_equal @half_year.length, 181
    assert_equal @one_month.length, 31
    assert_equal @two_month.length, 62
    assert_equal @one_day.length, 1
  end
  
  def test_step
    count = 0
    @half_year.step {|d| count += 1 }
    assert_equal count, 181
    count = 0
    @one_month.step {|d| count += 1 }
    assert_equal count, 31
    count = 0
    @two_month.step {|d| count += 1 }
    assert_equal count, 62
    count = 0
    @one_day.step {|d| count += 1 }
    assert_equal count, 1
  end
  
end