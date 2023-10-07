require 'sinatra'
require 'json'
require 'pry'
require 'active_model'

# we will mock having a state or database for the dev server by setting a global variable.
# Never use a global variable in a production server.
$home = {}

# Ruby class that includes validations from ActiveRecord.
# This will reprresent our home resource as a ruby object.
class Home
  # ActiveModel is part of Ruby on Rails used as an ORM.
  # It has a module within that provides validations.
  # The prod Terratowns server is rails and uses a similar validation.
  # https://guides.rubyonrails.org/active_model_basics.html
  include ActiveModel::Validations
  # Create virtual attributes to be stored in the object.
  # This will set a getter and setter
  # eg. home = new Home()
  # home.twon = 'hello' -- setter
  # home.town() -- getter
  attr_accessor :town, :name, :description, :domain_name, :content_version

  validates :town, presence: , inclusion: { in: [
  'cooker-cove',
  'melomaniac-mansion',
  'video-valley',
  'the-nomad-pad',
  'gamers-grotto'
  ]}
  # visible to all users
  validates :name, presence: true
  # visible to all users
  validates :description, presence: true
  # lock this down to only be from cloudfront using regex
  validates :domain_name, 
    format: { with: /\.cloudfront\.net\z/, message: "domain must be from .cloudfront.net" }
    # uniqueness: true, 
  # content-version has to be an integer
  # we will ensure it is an incremental version in the controller
  validates :content_version, numericality: { only_integer: true }
end


# extending a class from Sinatra::Base to enable this
# generic class utilize the Sinatra web framework
class TerraTownsMockServer < Sinatra::Base

  def error code, message
    halt code, {'Content-Type' => 'application/json'}, {err: message}.to_json
  end

  def error_json json
    halt code, {'Content-Type' => 'application/json'}, json
  end

  def ensure_correct_headings
    unless request.env["CONTENT_TYPE"] == "application/json"
      error 415, "expected Content_type header to be application/json"
    end

    unless request.env["HTTP_ACCEPT"] == "application/json"
      error 406, "expected Accept header to be application/json"
    end
  end

  # return hardcoded access token
  def x_access_code
    return '9b49b3fb-b8e9-483c-b703-97ba88eef8e0'
  end

  def x_user_uuid
    return '00522451-153b-4131-ab48-43e132f662cc'
  end

  def find_user_by_bearer_token
    auth_header = request.env["HTTP_AUTHORIZATION"]
    # check if AUTH header exists 
    if auth_header.nil? || !auth_header.start_with?("Bearer ")
      error 401, "a1000 Failed to authenicate, bearer token invalid and/or teacherseat_user_uuid invalid"
    end

    # Check if token matches the one in our database.
    # return an error if it does not match
    code = auth_header.split("Bearer ")[1]
    if code != x_access_code
      error 401, "a1001 Failed to authenicate, bearer token invalid and/or teacherseat_user_uuid invalid"
    end

    # check if user_uuid was present in payload json
    if params['user_uuid'].nil?
      error 401, "a1002 Failed to authenicate, bearer token invalid and/or teacherseat_user_uuid invalid"
    end

    # access token and user_uuid should be matching for user
    unless code == x_access_code && params['user_uuid'] == x_user_uuid
      error 401, "a1003 Failed to authenicate, bearer token invalid and/or teacherseat_user_uuid invalid"
    end
  end

  # CREATE
  post '/api/u/:user_uuid/homes' do
    ensure_correct_headings
    find_user_by_bearer_token
    # puts prints to the terminal similar to print in python
    puts "# create - POST /api/homes"

    # begin/rescue similar to try/catch block in python
    begin
      # Sinatra does not automatically parse json bodies as params
      # we need to manually parse it
      payload = JSON.parse(request.body.read)
    rescue JSON::ParserError
      halt 422, "Malformed JSON"
    end

    # assign payload to variables
    # Validate payload data
    name = payload["name"]
    description = payload["description"]
    domain_name = payload["domain_name"]
    content_version = payload["content_version"]
    town = payload["town"]

    # print variables to console
    puts "name #{name}"
    puts "description #{description}"
    puts "domain_name #{domain_name}"
    puts "content_version #{content_version}"
    puts "town #{town}"

    # Create new home model and set attribrutes
    home = Home.new
    home.town = town
    home.name = name
    home.description = description
    home.domain_name = domain_name
    home.content_version = content_version
    
    # unless validation checks pass otherwise return error message as json
    unless home.valid?
      error 422, home.errors.messages.to_json
    end

    # generate uuid at random
    uuid = SecureRandom.uuid
    puts "uuid #{uuid}"
    # mock save data to mock database
    $home = {
      uuid: uuid,
      name: name,
      town: town,
      description: description,
      domain_name: domain_name,
      content_version: content_version
    }

    # return uuid as json
    return { uuid: uuid }.to_json
  end

  # READ
  get '/api/u/:user_uuid/homes/:uuid' do
    ensure_correct_headings
    find_user_by_bearer_token
    puts "# read - GET /api/homes/:uuid"

    # checks for house limit

    content_type :json
    # does home uuid match uuid in mock database
    if params[:uuid] == $home[:uuid]
      return $home.to_json
    else
      error 404, "failed to find home with provided uuid and bearer token"
    end
  end

  # UPDATE
  put '/api/u/:user_uuid/homes/:uuid' do
    ensure_correct_headings
    find_user_by_bearer_token
    puts "# update - PUT /api/homes/:uuid"
    begin
      # Parse JSON payload from the request body
      payload = JSON.parse(request.body.read)
    rescue JSON::ParserError
      halt 422, "Malformed JSON"
    end

    # Validate payload data
    name = payload["name"]
    description = payload["description"]
    #domain_name = payload["domain_name"]
    content_version = payload["content_version"]

    unless params[:uuid] == $home[:uuid]
      error 404, "failed to find home with provided uuid and bearer token"
    end

    home = Home.new
    home.town = $home[:town]
    home.domain_name = $home[:domain_name]
    home.name = name
    home.description = description
    home.content_version = content_version

    unless home.valid?
      error 422, home.errors.messages.to_json
    end

    return { uuid: params[:uuid] }.to_json
  end

  # DELETE
  delete '/api/u/:user_uuid/homes/:uuid' do
    ensure_correct_headings
    find_user_by_bearer_token
    puts "# delete - DELETE /api/homes/:uuid"
    content_type :json

    if params[:uuid] != $home[:uuid]
      error 404, "failed to find home with provided uuid and bearer token"
    end

    # delete from mock database
    uuid = $home[:uuid]
    $home = {}
    { uuid: uuid }.to_json
  end
end


# This will run the server
TerraTownsMockServer.run!