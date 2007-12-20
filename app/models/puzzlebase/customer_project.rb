class Puzzlebase::CustomerProject < Puzzlebase::Base  
  
  set_table_name 'TBL_CUSTOMER_PROJECT'.downcase 
    
  belongs_to :customer,
             :foreign_key => 'FK_CUSTOMER'.downcase
  belongs_to :project,
             :foreign_key => 'FK_PROJECT'.downcase
             
  MAPS_TO = ::Project         
  MAPPINGS = {:shortname => :S_PROJECT.to_s.downcase.to_sym,
              :name      => :S_DESCRIPTION.to_s.downcase.to_sym } 
  FIND_OPTIONS = {:conditions => ["B_SYNCTOPUZZLETIME = 't'".downcase]}               
    
  # Synchronizes the Projects and the Customers.
  def self.synchronize
    resetErrors
    Puzzlebase::Unit.updateAll
    Puzzlebase::Customer.updateAll
    updateAll    
    errors
  end     
    
  protected  
  
  def self.updateLocal(original)
    success = super
    Puzzlebase::Project.updateChildren original.project, findLocal(original) if success
  end

  # TODO upcase original method names for MySQL
  def self.localFindOptions(original)
      { :include => :client,
        :conditions => ["projects.shortname = ? AND clients.shortname = ?", 
                        original.project.s_project, original.customer.s_customer] }
  end
  
  # TODO upcase original method names for MySQL
  def self.setReference(local, original)
    client = ::Client.find_by_shortname(original.customer.s_customer)
    local.client_id = client.id if client
  end
  
  def self.updateAttributes(local, original)
    super local, original.project
  end
end