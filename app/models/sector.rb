# encoding: utf-8
# == Schema Information
#
# Table name: sectors
#
#  id     :integer          not null, primary key
#  name   :string           not null
#  active :boolean          default(TRUE), not null
#

class Sector < ActiveRecord::Base

  has_many :clients, dependent: :nullify

  scope :list, -> { order(:name) }

  validates_by_schema
  validates :name, uniqueness: true

  def to_s
    name
  end

end