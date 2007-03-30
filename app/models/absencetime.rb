class Absencetime < Worktime
  def account
    absence
  end
  
  def account_id
    absence_id
  end
  
  def account_id=(value)
    self.absence_id = value
  end
  
  def absence?
    true
  end
  
  def self.account_label
    'Absenz'
  end
  
  def self.label
    'Absenz'
  end
  
end