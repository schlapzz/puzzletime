class EmployeesPlanningGraph

  include PlanningHelper
  
  attr_reader :period
  attr_reader :employees
  
  def initialize(employees, period = nil)
    @employees = employees.sort
    period ||= Period.currentMonth
    @actual_period = period
    @period = extend_to_weeks period
    @cache = Hash.new
    @colorMap = AccountColorMapper.new
    
    @employees.each do |employee|
      @cache[employee] = EmployeePlanningGraph.new(employee, period)
    end
    
  end
  
  def graphFor(user)
    @cache[user]
  end

  def colorFor(project)
    @colorMap[project]
  end
  
end